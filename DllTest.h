#pragma once

#include <stdio.h>
#include <stdlib.h>
//#include <string.h>
//#include <iostream>



struct structA 
{
	int 	Key;
	char 	buff[1024];
};

#ifdef _WIN32
	#define DLLIMPORT __declspec(dllexport)
	DLLIMPORT int add(int x, int y);
#else
	int Add(int x, int y);
	int Sub(int x, int y);
#endif