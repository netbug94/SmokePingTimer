# SmokePing App has been overloading and crashing in the latest versions of libreNMS integration, This app runs systemctl Smokeping every week to keep it from happening.

# How to Add an Exemption to sudo

Follow these steps to allow a specific user to execute a command without a password prompt using `sudo`.

### 1. Identify the Full Path
First, you need to determine the full path to the executable or script for which you want to grant sudo access without a password. For example:
```
/home/olracnai/FileName
```

### 2. Edit the sudoers File Using `visudo`
To safely edit the sudoers file, use the `visudo` command. This command checks for syntax errors before saving, which helps prevent configuration errors that could potentially lock you out of the system.

- Open your terminal.
- Enter the following command:
  ```bash
  sudo visudo
  ```
- Scroll to the end of the file and add the following line:
  ```
  olracnai ALL=(ALL) NOPASSWD: /home/olracnai/FileName
  ```
  Replace `olracnai` with your actual username and `/home/olracnai/FileName` with the full path to your executable.

### 3. Test the Exemption
After adding the exemption, you should test to ensure that it works as expected without requiring a password.

- Run the following command in your terminal:
  ```bash
  sudo -n /home/olracnai/FileName
  ```
  Make sure to replace `/home/olracnai/FileName` with the actual path to your file.

If set up correctly, this command should execute without prompting you for a password. If it fails, it will typically provide an error message indicating that a password is needed, which suggests that the exemption was not properly set in the sudoers file.

--- 
