# GoShark
A wireshark CLI written in go!

CURRENTLY MUST BE RUN AS ROOT

this is because libpcap is protected at the kernel level and to avoid permissions mishaps, you must be root to capture packets

Basic Premise
- CLI
- specifiy an interface
- capture packets
- Send them to ~a MongoDB on~ GCP
- analyze the packets for insecurities
- alert the user of insecure connections

TODO:
- ~Be able to list interfaces~
- ~Take in user input to specifiy interface~
- ~start capturing packets in background while user does stuff~
- every 5 min, send packets to GCP to be analyzed

ON GCP
- receive packets
- use REGEX to find insecure connections, and filter those out immediately (ie, unsecured telnet, flag)
- write alerts to latex documents (then compile)
- SCP file to host

- Host then has pdf with alerts
