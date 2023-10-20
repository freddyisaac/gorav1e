#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "rav1e.h"
#include "glue.h"

RaChromaticityPointArray* new_chromacity_point_array(int n)
{
	RaChromaticityPointArray *car;
	car = (RaChromaticityPointArray*) calloc(1, sizeof(RaChromaticityPointArray));
	car->num = 0;
	car->max = n;
	car->array = (RaChromaticityPoint*) calloc(n, sizeof(RaChromaticityPoint));
	return car;
}

void add_chromacity_point(RaChromaticityPointArray* carr, RaChromaticityPoint pt)
{
	if (!carr) return;
	carr->num++;
	if (carr->num == carr->max)
	{
		carr->max += 4;
		carr->array = (RaChromaticityPoint*) realloc((void*) carr->array, carr->max * sizeof(RaChromaticityPoint));
	}
	memcpy((void*) &carr->array[carr->num-1], (const void*) &pt, sizeof(RaChromaticityPoint));
} 

void free_chromacity_point_array(RaChromaticityPointArray *carr)
{
	if (!carr) return;
	if (carr->array) free(carr->array);
	free(carr);
}
