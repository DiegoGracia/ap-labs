//Diego Armando Gracia Hinojosa
//A01229716


#include <stdio.h>
#include <stdarg.h>

#define RESET			0
#define BRIGHT 		1
#define DIM				2
#define UNDERLINE 3
#define BLINK			4
#define REVERSE		7
#define HIDDEN		8

#define BLACK 		0
#define RED				1
#define GREEN			2
#define YELLOW		3
#define BLUE			4
#define MAGENTA		5
#define CYAN			6
#define	WHITE			7

void
textcolor (int attr, int fg, int bg)
{
  char command[13];
  sprintf (command, "%c[%d;%d;%dm", 0x1B, attr, fg + 30, bg + 40);
  printf ("%s", command);
}

int
infof (const char *format, ...)
{
  va_list arg;
  int done;

  textcolor (BRIGHT, BLUE, HIDDEN);
  printf ("INFO: ");
  textcolor (RESET, WHITE, HIDDEN);

  va_start (arg, format);
  done = vprintf (format, arg);
  va_end (arg);
  return done;
}

int
warnf (const char *format, ...)
{
  va_list arg;
  int done;

  textcolor (BRIGHT, YELLOW, HIDDEN);
  printf ("WARN: ");
  textcolor (RESET, WHITE, HIDDEN);

  va_start (arg, format);
  done = vprintf (format, arg);
  va_end (arg);
  return done;
}

int
errorf (const char *format, ...)
{
  va_list arg;
  int done;

  textcolor (BRIGHT, RED, HIDDEN);
  printf ("ERROR: ");
  textcolor (RESET, WHITE, HIDDEN);

  va_start (arg, format);
  done = vprintf (format, arg);
  va_end (arg);
  return done;
}

int
panicf (const char *format, ...)
{
  va_list arg;
  int done;

  textcolor (BRIGHT, WHITE, RED);
  printf ("PANIC:");
  textcolor (RESET, WHITE, HIDDEN);
  printf (" ");

  va_start (arg, format);
  done = vprintf (format, arg);
  va_end (arg);
  return done;
}
