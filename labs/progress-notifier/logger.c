#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <signal.h>
#include <sys/types.h>
#include <unistd.h>
#include <string.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <syslog.h>
#include "logger.h"

#define RESET 0
#define BRIGHT 1
#define DIM	2
#define UNDERLINE 3
#define BLINK	4
#define REVERSE 7
#define HIDDEN 8

#define BLACK 0
#define RED 1
#define GREEN 2
#define YELLOW 3
#define BLUE 4
#define MAGENTA 5
#define CYAN 6
#define	WHITE 7

#define STDOUT 0
#define SYSLOG 1

int log = 0;

/*
  Using the colors references that the profesor give us in a link,
  I find textcolor function and #define that are made by Pradeep Padala.
*/
void textcolor(int attr, int fg, int bg) {
  char command[13];
	/* Command is the control command to the terminal */
	sprintf(command, "%c[%d;%d;%dm", 0x1B, attr, fg + 30, bg + 40);
	printf("%s", command);
}

int initLogger(char *logType) {
    if (strcmp(logType, "syslog") == 0) {
        log = SYSLOG;
        printf("\nInitializing Logger on: %s\n", logType);
    }
    else if (strcmp(logType, "stdout") == 0) {
        log = STDOUT;
        printf("\nInitializing Logger on: %s\n", logType);
    }
    else {
        printf("\nInvalid Log\n");
    }
    return 0;
}

/*
  I find this information about Variable Argument Lists, that would help me to create the solution
  https://www.cprogramming.com/tutorial/c/lesson17.html
*/
int infof(const char *format, ...) {
  int flag;
  va_list arg;
  if (log == STDOUT) {
      textcolor (BRIGHT, WHITE, HIDDEN);
      printf ("INFO: ");
      textcolor (RESET, WHITE, HIDDEN);
      va_start (arg, format);
      flag = vprintf (format, arg);
      va_end (arg);
      return flag;
  }
  else {
      openlog("syslog", LOG_NDELAY, LOG_DAEMON);
      vsyslog(LOG_INFO, format, arg);
      closelog();
      va_end(arg);
      return 1;
  }
}

int warnf(const char *format, ...) {
  int flag;
  va_list arg;

  if (log == STDOUT) {
      textcolor (BRIGHT, YELLOW, HIDDEN);
      printf ("\nWARNING: ");
      textcolor (RESET, YELLOW, HIDDEN);
      va_start (arg, format);
      flag = vprintf (format, arg);
      va_end (arg);
      return flag;
  }
  else {
      openlog("syslog", LOG_NDELAY, LOG_DAEMON);
      vsyslog(LOG_INFO, format, arg);
      closelog();
      va_end(arg);
      return 1;
  }
}

int errorf(const char *format, ...) {
  int flag;
  va_list arg;
  if (log == STDOUT) {
      textcolor (BRIGHT, RED, HIDDEN);
      printf ("\nERROR: ");
      textcolor (RESET, RED, HIDDEN);
      va_start (arg, format);
      flag = vprintf (format, arg);
      va_end (arg);
      return flag;
  }
  else {
      openlog("syslog", LOG_NDELAY, LOG_DAEMON);
      vsyslog(LOG_INFO, format, arg);
      closelog();
      va_end(arg);
      return 1;
  }
}

int panicf(const char *format, ...) {
  int flag;
  va_list arg;
  va_start (arg, format);
  if (log == STDOUT) {
      textcolor (BRIGHT, MAGENTA, HIDDEN);
      printf ("\nsPANIC: ");
      textcolor (RESET, MAGENTA, HIDDEN);
      va_start (arg, format);
      flag = vprintf (format, arg);
      va_end (arg);
      return flag;
  }
  else {
      openlog("syslog", LOG_NDELAY, LOG_DAEMON);
      vsyslog(LOG_INFO, format, arg);
      closelog();
      va_end(arg);
      return 1;
  }
}
