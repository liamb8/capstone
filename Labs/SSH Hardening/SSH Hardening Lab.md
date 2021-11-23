# Work In Progress



# SSH Hardening Lab



To start off open the SSH config file

`sudo nano /etc/ssh/sshd_config` 



## Disabling Password-Based Authentication



To only use public key authentication to login to a user account you first need to generate an ssh key from the host machine in which your vm in stored on or another vm that runs Ubuntu or Linux. To generate an ssh key type in the command line `ssh-keygen` this will generate the ssh key. After this you need to copy the ssh key to the vm where you will be sshing into. To do this type `ssh-copy-id cyberlabs@ip-address` in the command line. 



After doing those steps you will need to edit the SSH config file in order for SSH to authenticate with a public key. Bellow are the 2 lines that you will need to uncomment or add within   the SSH config file on the Ubuntu vm. You will want to change `PubkeyAuthentication` from no to yes and `PasswordAuthentication` from yes to no.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/PasswordAuth.JPG)



## Disabling Empty Passwords



To disable empty passwords add or uncomment the line `PermitEmptyPasswords yes` and change it to no. This will prevent anyone from trying to login without using a password.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/PasswordAuth.JPG)



## Forbidding Root Login



To disable root login uncomment the `PermitRootLogin yes` and change it to no. This will disable anyone from attempting to login under the `root` account.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/PermitRootLogin.JPG)



## Using SSH Protocol 2



By default SSH uses Protocol 1 to change this add `Protocol 2` in the SSH config file. The will allow SSH to run on Protocol 2 instead of the default Protocol 1.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/Protocol.JPG)



## Setting a Session Timeout



To set a session timeout uncomment the line `ClientAliveInterval` and change the value that was set there to `300`. This will set a session timeout of 300 seconds or 5 minutes.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/ClientAliveInt.JPG)



## Allow Specific Users Access to SSH



`AllowUsers john`



## Limiting Number of Authentication Attempts



To limit the number of authentication attempts when trying to SSH into an account uncomment the line `MaxAuthTries ` and set the value to `3`. This will set the number of authentication attempts allowed to 3 and after 3 failed authentication attempts the ssh session will close.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/MaxAuth.JPG)

