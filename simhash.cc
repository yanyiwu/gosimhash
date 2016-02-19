extern "C" {
#include "simhash.h"
}

#include "simhash/Simhasher.hpp"

Simhasher NewSimhasher(const char* dict_path, 
      const char* model_path,
      const char* idf_path,
      const char* stop_words_path) {
  Simhasher x = (Simhasher)(new simhash::Simhasher(dict_path, model_path, idf_path, stop_words_path));
  return x;
}

void FreeSimhasher(Simhasher x) {
  delete (simhash::Simhasher*)x;
}

uint64_t MakeSimhash(Simhasher x, const char* sentence, int top_n) {
  uint64_t result = 0;
  ((simhash::Simhasher*)x)->make(sentence, top_n, result);
  return result;
}
