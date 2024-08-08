# Evilbytecode-Anti-VM
- Detect if user is running inside a virtual machine.
> [!NOTE]
> Only works on x86 and x86-64.

The code measures the average CPU cycles required to execute the ```CPUID``` instruction. It compares this value to a predefined threshold. If the measured cycles exceed the threshold, the code assumes it's running inside a virtual machine (VM) due to the higher overhead typically associated with VMs.

# Credits
- him for the idea lol.
- https://github.com/zwinplayer64

# PoC:
![image](https://github.com/user-attachments/assets/5bb98ad9-88c4-47f4-9d2d-7b6168d28eca)
