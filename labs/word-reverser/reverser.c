//Diego Armando Gracia Hinojosa
//A01229716


#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static int n;

void reverse(int lenght, char word[]) {
	n = 0;
	int i;
	char tmp;
	for(i = 0; i < lenght/2; i++) {
		tmp = word[i];
		word[i] = word[lenght - i - 1];
		word[lenght - i - 1] = tmp;
	}
	printf("%s\n", word);
}

int main(int argc, char **argv){
	char c;
	char word[100];
	n = 0;
	while ((c = getchar()) != EOF) {
		word[n++] = c;
		if(c == '\n') {
			reverse(n, word);
			memset(word, 0, 100);
		}
	}
	return 0;
}
