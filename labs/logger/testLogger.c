

int infof(const char *format, ...);
int warnf(const char *format, ...);
int errorf(const char *format, ...);
int panicf(const char *format, ...);

int main()
{
infof("This is information test\n");
warnf("This is a %s test\n", "warn");
errorf("This is an %s %s\n", "error", "test");
panicf("This is a panic test\n");
return 0;
}

