#ifndef CPPJIEBA_MEMSTREAM_H
#define CPPJIEBA_MEMSTREAM_H

class memstream : public std::istream {
  struct membuf : std::streambuf {
    membuf(const uint8_t *p, size_t l) {
      setg((char *)p, (char *)p, (char *)p + l);
    }
  };
  membuf _buffer;

 public:
  memstream(const uint8_t *p, size_t l)
      : std::istream(&_buffer), _buffer(p, l) {
    rdbuf(&_buffer);
  }
};

#endif
