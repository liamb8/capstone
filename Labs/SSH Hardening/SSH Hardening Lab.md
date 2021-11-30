# SSH Hardening Lab

**[Virtual Machine Download Link](https://drive.google.com/file/d/1U4osXUWvy-8-l_oP_Cs8HiVIPt3mfiDx/view?usp=sharing)**

------

Before starting the lab, make sure to **download the ZIP file provided**. After downloading, **unzip the file** and place the folder wherever you find appropriate. The instructions below assume that you're using VMware Workstation Client, although you can still follow them if you're using other VM clients.

Go to the **File** tab in the upper left-hand corner of the VM clients window and click on it, a drop-down menu will show up. From there, click on **Open**. After that, a file explorer will open navigate to where you put the VM you unzipped. Click on the `SSHHardeningLab.vmx` image and it should show up in your client. If this doesn't work click on the `SSHHardeningLab.ovf` and it will bring up an import menu. 

Once you have done this click on the **Import** button and the VM should show up. After this right click on the VM and click on **settings**, this is where you can change the settings for your VM. Below is an image of what the settings of the VM should be. After checking that all the settings are the same click **Ok** and then right-click on the VM and choose power and then click on **Start up guest** to power on the VM.



### User Accounts:

Below is the user account which will be used for this lab.

- John - Administrator
  - User Login: `john`
  - Password: `Fallentrees4`



## Disabling Password-Based Authentication

Open the SSH config file

`sudo nano /etc/ssh/sshd_config` 

Instead of using a password to log in to an account using an ssh key is considered more secure. By using an ssh key it helps to protect an account better than using a password to log in. With a password, anyone can log in to your account as long as they know the password set. With ssh keys, this is a different story as only the machine which copied its ssh key to the machine that it wants to ssh into can log in to it. This means that only those with that key can log in unlike with a password where anyone can log in as long as they know it.

To only use public-key authentication to log in to a user account, you first need to generate an ssh key from the host machine where your vm is stored on or another vm that runs Ubuntu or Linux. To generate an ssh key type in the command line `ssh-keygen`. After this, you need to copy the ssh key to the vm where you will be sshing into. To do this type `ssh-copy-id john@ip-address` in the command line. 

After doing those steps you will need to edit the SSH config file in order for SSH to authenticate with a public key. Bellow are the 2 lines that you will need to uncomment within   the SSH config file on the Ubuntu vm. You will want to change `PubkeyAuthentication` from no to **yes** and `PasswordAuthentication` from yes to **no**.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/PasswordAuth.JPG)



## Disabling Empty Passwords

Why disable empty passwords? Sometimes people may think that having an empty password is convenient and that they won't have to remember what they used. Although it may be convenient it is also extremely insecure. By leaving a password empty anyone can log in to the account.

To disable empty passwords add or uncomment the line `PermitEmptyPasswords yes` and change it to no. This will prevent anyone from trying to log in without using a password.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/PasswordAuth.JPG)



## Forbidding Root Login

By forbidding root login you're restricting access to anyone who wants to ssh into the root account. This is important so that no intruders can gain root-level access. For example, if a hacker were to get root-level access it would be extremely detrimental to the computer and the company itself. Especially when that computer could be connected to other networks or services.

To disable root login uncomment the `PermitRootLogin yes` and change it to no. This will disable anyone from attempting to log in under the `root` account.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/PermitRootLogin.JPG)



## Using SSH Protocol 2

An SSH server can use two different protocols like Protocol 1 and Protocol 2. The reason why it's important to use Protocol 2 for ssh is that it implements more advanced security features unlike Protocol 1. 

By default, SSH uses Protocol 1 to change this add `Protocol 2` in the SSH config file. The will, allow SSH to run on Protocol 2 instead of the default Protocol 1.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/Protocol.JPG)



## Setting a Session Timeout

Sometimes people leave their computers unattended for long periods. It's important to set a session timeout for this reason as an intruder could gain access to your system if you leave a session open without closing it. For example, you may leave your desk at work and you have an ssh session opened someone could use your computer during that time while also gaining access to the system that you're connected to. This could also happen if your computer was breached itself as leaving a session open without a time limit to disconnect could allow an intruder to use that system for malicious purposes.

To set a session timeout, uncomment the line `ClientAliveInterval` and change the value that was set there to `300`. This will set a session timeout of 300 seconds or 5 minutes.

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/ClientAliveInt.JPG)



## Limiting Number of Authentication Attempts

When a user tries to log in to a system and is unable to authenticate themselves once they will continue to try to log in. The user will keep trying to log in until they can successfully authenticate themselves and gain access to the system. This is considered highly insecure as a hacker could launch a Brute Force Attack (an attack that repeatedly attempts to guess a password until a match is found). This is why it's important to limit the number of authentication attempts so it can deter these types of attacks. 

To limit the number of authentication attempts when trying to SSH into an account uncomment the line `MaxAuthTries ` and set the value to `3`. This will set the number of authentication attempts allowed to 3 and after 3 failed authentication attempts the ssh session will close.

Once finished with all the previous steps, restart the ssh service for the changes made to take effect. You can do this by doing `sudo systemctl restart sshd`. 

![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/MaxAuth.JPG)

