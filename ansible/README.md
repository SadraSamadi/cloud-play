# VPS Preparation for Ansible

This guide explains all the necessary steps to prepare a fresh VPS so it can be managed with Ansible. Follow each step
carefully. All commands assume you are starting from a clean Ubuntu 22.04 LTS VPS.

---

## 1. Log in to the VPS as root

```bash
ssh root@VPS_ADDRESS
````

Replace `VPS_ADDRESS` with the VPS public address (host/ip).

---

## 2. Create and configure an Ansible user

Create a new user called `ansible`:

```bash
adduser ansible
```

Add the user to the `sudo` group:

```bash
usermod -aG sudo ansible
```

---

## 3. Add your SSH key for the Ansible user

```bash
mkdir -p /home/ansible/.ssh
cp ~/.ssh/authorized_keys /home/ansible/.ssh/
chown -R ansible:ansible /home/ansible/.ssh
chmod 700 /home/ansible/.ssh
chmod 600 /home/ansible/.ssh/authorized_keys
```

Test login as the new user:

```bash
exit
ssh ansible@VPS_ADDRESS
```

---

## 4. Enable passwordless sudo for the Ansible user

Edit the sudoers file:

```bash
sudo visudo
```

Add the following line at the bottom:

```
ansible ALL=(ALL) NOPASSWD:ALL
```

Test:

```bash
sudo whoami
```

Expected output: `root` without a password prompt.

---

## 5. Check Python version

Ansible requires Python on the VPS:

```bash
python3 --version
```

If missing, install:

```bash
sudo apt update
sudo apt install -y python3 python3-apt
```

---

## 6. Update and clean up the system

```bash
sudo apt update && sudo apt upgrade -y
sudo apt full-upgrade -y
sudo apt autoremove -y
sudo apt clean -y
```

---

## 7. Reboot the system

```bash
sudo reboot
```

Wait for the VPS to come back online, then reconnect:

```bash
ssh ansible@VPS_ADDRESS
```

---

## 8. Add the VPS to Ansible inventory

Update inventory file and add the new VPS (`inventory.yaml`):

```yaml
all:
  hosts:
    vps:
      ansible_host: VPS_ADDRESS
      ansible_port: 22
      ansible_user: ansible
```

---

## 9. Create an SSH agent for the terminal

```bash
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_rsa
```

Replace `id_rsa` with your private key if different.

---

## 10. Test Ansible connectivity

```bash
ansible -i inventory.yaml vps -m ping
```

Expected output:

```
vps | SUCCESS => pong
```

---

## 11. Run Ansible playbooks

From the project root, run any playbook:

```bash
ansible-playbook -i inventory.yaml playbooks/<playbook>.yml
```

---

✅ After completing these steps, your VPS is fully prepared for Ansible-managed automation.

