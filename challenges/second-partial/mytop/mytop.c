#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <string.h>
#include <dirent.h>
#include <ctype.h>
#include <time.h>
#include <signal.h>

#define DIRECTORYSIZE 512
#define numOfBytes 512

void clear();
void fileWriter(char *tempFile);
void printerToFile(int fd);
void interrupStatusFile(int signal);
char *exploreStatus(char *path, int openFiles);
void exitOnForce(int signal);
void clean(void);
int openFileExplorer(char *path);

int main() {
	if (signal(SIGINT, interrupStatusFile) == SIG_ERR)
		printf("Signal map wasn't successful");
	if (signal(SIGQUIT, exitOnForce) == SIG_ERR)
		printf("Signal map wasn't successful");
	clear();
	while (1) {
		clear();
		fileWriter(".temp");
	}
	atexit(clean);
	return 0;
}

void clear() {
	printf("\e[1;1H\e[2aux");
}

void fileWriter(char *tempFile) {
	DIR *dir;
	struct dirent *ent;
	char *procDirectory, *resToFile;
	int fd;
	procDirectory = calloc(DIRECTORYSIZE, sizeof(char));
	resToFile = calloc(numOfBytes, sizeof(char));
	fd = open(tempFile, O_RDWR | O_CREAT, S_IRWXO);
	if (fd == -1) {
		close(fd);
		printf("Couldn't create file");
		exit(EXIT_FAILURE);
	}
	sprintf(resToFile, "|  %-9s|  %-9s|  %-41s|  %-15s|  %-15s|  %-9s|  %-12s|\n",
		"PID", "Parent", "Name", "State", "Memory", "# Threads", "Open Files");
	write(fd, resToFile, strlen(resToFile));
	sprintf(resToFile,
		"|-----------|-----------|-------------------------------------------|-----------------|-----------------|-----------|--------------|\n");
	write(fd, resToFile, strlen(resToFile));
	strcpy(procDirectory, "/proc");
	dir = opendir(procDirectory);
	if (dir != NULL) {
		while ((ent = readdir(dir)) != NULL) {
			if (isdigit(ent->d_name[0])) {
				strcpy(procDirectory, "/proc/");
				strcat(procDirectory, ent->d_name);
				sprintf(resToFile, "%s", exploreStatus(procDirectory, openFileExplorer(procDirectory)));
				if (strlen(resToFile) > 100)
					write(fd, resToFile, strlen(resToFile));
				memset(procDirectory, '\0', DIRECTORYSIZE);
				memset(resToFile, '\0', numOfBytes);
			}
		}
	}
	closedir(dir);
	printerToFile(fd);
	close(fd);
}

char *exploreStatus(char *path, int openFiles){
	char *statusFile, *buffer, *child, *status, *value;
	char *pid, *ppid, *name, *state, *memory, *threads;
	int fd, flag, auxump, aux = 0;
	double mbs = 0;
	statusFile = calloc(DIRECTORYSIZE + 7, sizeof(char));
	buffer = calloc(numOfBytes, sizeof(char));
	child = calloc(8, sizeof(char));
	status = calloc(13, sizeof(char));
	pid = calloc(6, sizeof(char));
	ppid = calloc(6, sizeof(char));
	name = calloc(256, sizeof(char));
	state = calloc(13, sizeof(char));
	memory = calloc(20, sizeof(char));
	threads = calloc(2, sizeof(char));
	value = calloc(256, sizeof(char));
	strcpy(statusFile, path);
	strcat(statusFile, "/status");
	fd = open(statusFile, O_RDONLY);
	if (fd == -1) {
		close(fd);
		return "";
	}
	while (read(fd, buffer, numOfBytes - 1) > 0) {
		flag = 0;
		for (int i = 0; buffer[i] != '\0'; i++) {
			if (buffer[i] != '\n') {
				if (buffer[i] != ':' && flag == 0) {
					child[i] = buffer[i];
				}
                else if (buffer[i] == ':') {
					flag = 1;
				}
				if (flag) {
					if (buffer[i] != ':' && buffer[i] != '\t') {
						if (strcmp(child, "Pid") == 0
						    || strcmp(child, "PPid") == 0
						    || strcmp(child, "Name") == 0
						    || strcmp(child, "State") == 0
						    || strcmp(child, "Threads") == 0
						    || strcmp(child, "VmSize") == 0) {
							  status[aux] = buffer[i];
							  aux++;
						}
					}
					if (buffer[i + 1] == '\n') {
						if (strcmp(child, "Pid") == 0) {
							strcpy(pid, status);
						}
                        else if (strcmp(child, "PPid") == 0) {
							strcpy(ppid, status);
						}
                        else if (strcmp(child, "Name") == 0) {
							strcpy(name, status);
						}
                        else if (strcmp(child, "State") == 0) {
							strcpy(state, status);
						}
                        else if (strcmp(child, "VmSize") == 0) {
							mbs = (atoi(status) + .0) / 1000;
							sprintf(memory, "%.4fM", mbs);
						}
                        else if (strcmp(child, "Threads") == 0) {
							strcpy(threads, status);
						}
						flag = 0;
						memset(status, '\0', 13);
					}
				}
			}
            else {
				auxump = (i * sizeof(char) - strlen(buffer) *sizeof(char)) + 1;
				aux = 0;
				lseek(fd, auxump, SEEK_CUR);
				memset(child, '\0', 8);
				memset(buffer, '\0', numOfBytes);
			}
		}
	}
	close(fd);
	if (memory[0] == '\0')
		strcpy(memory, "0.000M");
	sprintf(value, "|  %-9s|  %-9s|  %-41s|  %-15s|  %-15s|  %-9s|  %-12d|\n", pid,
		ppid, name, state, memory, threads, openFiles);
	return value;
}

int openFileExplorer(char *path) {
	DIR *dir;
	struct dirent *ent;
	char *procDirectory;
	int dirs = 0;
	procDirectory = calloc(DIRECTORYSIZE, sizeof(char));
	strcpy(procDirectory, path);
	strcat(procDirectory, "/fd");
	dir = opendir(procDirectory);
	if (dir != NULL) {
		while ((ent = readdir(dir)) != NULL) {
			if (isdigit(ent->d_name[0])) {
				dirs++;
			}
		}
	}
	closedir(dir);
	return dirs;
}

void printerToFile(int fd) {
	char *buffer;
	buffer = calloc(numOfBytes, sizeof(char));
	lseek(fd, 0, SEEK_SET);
	while (read(fd, buffer, numOfBytes - 1) > 0) {
		printf("%.*s", (int)strlen(buffer), buffer);
		memset(buffer, '\0', numOfBytes);
	}
}

void interrupStatusFile(int signal) {
	char date[11], title[25];
	time_t t = time(NULL);
	struct tm tm = *localtime(&t);
	sprintf(date, "%d-%d-%d", tm.tm_mday, tm.tm_mon + 1, tm.tm_year + 1900);
	strcpy(title, "mytop-status-");
	strcat(title, date);
	fileWriter(title);
}
