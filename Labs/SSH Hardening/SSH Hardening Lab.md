# Work In Progress



# SSH Hardening Lab





`sudo nano /etc/ssh/sshd_config` 



## Disabling Password-Based Authentication



`ssh-keygen`

`ssh-copy-id cyberlabs@ip-address`

`PubkeyAuthentication yes`

`PasswordAuthentication no`



## Disabling Empty Passwords



`PermitEmptyPasswords no`



## Forbidding Root Login



`PermitRootLogin no`



## Using SSH Protocol 2



`Protocol 2`



## Setting a Session Timeout



`ClientAliveInterval 300`



## Allow Specific Users Access to SSH



`AllowUsers john`



## Limiting Number of Authentication Attempts



`MaxAuthTries 3`



