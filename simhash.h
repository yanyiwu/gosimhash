#ifndef GOSIMHASH_SIMHASH_H
#define GOSIMHASH_SIMHASH_H

#include <stdint.h>

typedef void* Simhasher;

Simhasher NewSimhasher(const char* dict_path, 
      const char* model_path,
      const char* idf_path,
      const char* stop_words_path);

void FreeSimhasher(Simhasher x);

uint64_t MakeSimhash(Simhasher x, const char* sentence, int top_n);

#endif // GOSIMHASH_SIMHASH_H
