#include <stdint.h>
#include <cpuid.h>  // For __get_cpuid function

// Function to read the timestamp counter
uint64_t read_tsc() {
    unsigned int low, high;
    __asm__ __volatile__("rdtsc" : "=a"(low), "=d"(high));
    return ((uint64_t)high << 32) | low;
}

// Function to perform CPUID instruction
void perform_cpuid(unsigned int *eax, unsigned int *ebx, unsigned int *ecx, unsigned int *edx) {
    __get_cpuid(0, eax, ebx, ecx, edx);
}
