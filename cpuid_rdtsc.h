#ifndef CPUID_RDTSCP_H
#define CPUID_RDTSCP_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

uint64_t read_tsc();
void perform_cpuid(unsigned int *eax, unsigned int *ebx, unsigned int *ecx, unsigned int *edx);

#ifdef __cplusplus
}
#endif

#endif // CPUID_RDTSCP_H
