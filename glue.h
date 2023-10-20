#ifndef GLUE_H
#define GLUE_H

typedef struct {
	int num;
	int max;
	RaChromaticityPoint *array;
} RaChromaticityPointArray;

extern RaChromaticityPointArray* new_chromacity_point_array(int);
extern void free_chromacity_point_array(RaChromaticityPointArray*);
extern void add_chromacity_point(RaChromaticityPointArray*, RaChromaticityPoint);

#endif //GLUE_H

