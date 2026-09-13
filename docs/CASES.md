# Initial Diagnostic Cases

1. **Connected to WiFi but No Internet**
   Diagnose why a network connection exists but internet access is unavailable by checking DNS, routing, DHCP, gateway reachability, and driver issues.

2. **WiFi Authenticates but Connection Fails**
   Detect WPA handshake failures, firmware issues, driver problems, or access-point compatibility issues preventing successful connection.

3. **Random WiFi Disconnects**
   Identify recurring wireless disconnections caused by driver resets, firmware crashes, power management, or hardware instability.

4. **Service Running but Not Reachable**
   Determine why a service reported as active cannot be accessed due to binding, firewall, routing, proxy, or namespace issues.

5. **High Load Average with Low CPU Usage**
   Explain situations where system load is high despite low CPU utilization by detecting blocked tasks, I/O waits, or filesystem stalls.

6. **Unexpected Reboot Analysis**
   Investigate previous boot logs to identify whether a reboot was caused by power loss, kernel panic, watchdog reset, or hardware failure.

7. **System Freeze Before Reboot**
   Perform post-mortem analysis of freezes caused by storage hangs, driver issues, blocked tasks, or kernel lockups.

8. **Thermal Throttling Detection**
   Detect CPU frequency throttling caused by overheating and explain its impact on system performance.

9. **Interrupt Storm Detection**
   Identify abnormal hardware interrupt activity from devices such as NICs, USB controllers, or drivers causing CPU overhead.

10. **Process Refuses to Exit**
    Explain why a process cannot be terminated due to D-state waits, stopped jobs, child processes, or kernel-level blocking.

11. **Disk Full Despite File Deletion**
    Detect processes holding deleted files open and explain why disk space has not been reclaimed.

12. **VPN Connected but No Traffic**
    Diagnose routing, DNS, MTU, or split-tunneling issues that prevent traffic from flowing through a connected VPN.

13. **Driver/Firmware Reset Loop**
    Detect repeated hardware or firmware resets affecting devices such as WiFi adapters, Bluetooth modules, GPUs, or NICs.
