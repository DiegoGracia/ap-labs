#include <stdio.h>
#include <signal.h>
#include <stdarg.h>
#include <string.h>
#include <syslog.h>
#include "logger.h"

#define RESET		0
#define BRIGHT 		1
#define DIM		2
#define UNDERLINE 	3
#define BLINK		4
#define REVERSE		7
#define HIDDEN		8

#define BLACK 		0
#define RED		1
#define GREEN		2
#define YELLOW		3
#define BLUE		4
#define MAGENTA		5
#define CYAN		6
#define	WHITE		7

int logFlag = 0;

void textcolor(int attr, int fg, int bg)
{	char command[13];

	/* Command is the control command to the terminal */
	sprintf(command, "%c[%d;%d;%dm", 0x1B, attr, fg + 30, bg + 40);
	printf("%s", command);
}

int initLogger(char *logType) {
	if((strcmp(logType, "")) == 0) {
		// Default
		logFlag = 0;
		//printf("Initializing Logger on: stdout\n");
	}
	else if((strcmp(logType, "stdout")) == 0) {
		// Stdout
		logFlag = 0;
	}
	else if((strcmp(logType, "syslog")) == 0) {
		// Syslog
		logFlag = 1;
	}
	else {
		logFlag = 0;
		printf("Not recognized log destination, going to default stdout\n");
	}
    printf("Initializing Logger on: %s\n", logType);
    return 0;
}

int infof(const char *format, ...) {
	va_list arg;
	va_start(arg, format);
	int done;
	// Normal text color
	textcolor(RESET, WHITE, BLACK);

	if(logFlag) {
		openlog("myLog", LOG_NDELAY, LOG_DAEMON);
		vsyslog(LOG_INFO, format, arg);
		closelog();
	}
	else {
		done = vfprintf(stdout, format, arg);
	}

	va_end(arg);
	return done;
}

int warnf(const char *format, ...) {
	va_list arg;
	va_start(arg, format);
        int done;
	// Warning text color
        textcolor(RESET, YELLOW, BLACK);

	if(logFlag) {
		openlog("myLog", LOG_NDELAY, LOG_DAEMON);
		vsyslog(LOG_WARNING, format, arg);
		closelog();
	}
	else {
        	done = vfprintf(stdout, format, arg);
	}

        va_end(arg);
	textcolor(RESET, WHITE, BLACK);
        return done;
}

int errorf(const char *format, ...) {
	va_list arg;
        va_start(arg, format);
	int done;
	// Error text color
        textcolor(BRIGHT, RED, BLACK);

	if(logFlag) {
		openlog("myLog", LOG_NDELAY, LOG_DAEMON);
		vsyslog(LOG_ERR, format, arg);
		closelog();
	}
	else {
		done = vfprintf(stdout, format, arg);
	}

        va_end(arg);
        textcolor(RESET, WHITE, BLACK);
        return done;
}

int panicf(const char *format, ...) {
	va_list arg;
	va_start(arg, format);
        int done;
        // Normal text color
        textcolor(RESET, WHITE, BLACK);

	if(logFlag) {
		openlog("myLog", LOG_NDELAY, LOG_DAEMON);
		vsyslog(LOG_EMERG, format, arg);
		closelog();
	}
	else {
		printf("PANIC MODE: ");
		done = vfprintf(stdout, format, arg);
	}

        va_end(arg);
        textcolor(RESET, WHITE, BLACK);

	raise(SIGABRT);
        return -1;
}
