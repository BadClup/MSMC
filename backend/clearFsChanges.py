import os
import subprocess


def list_containers():
    result = subprocess.run(['docker', 'ps', '-a', '--format', '{{.Names}}'], capture_output=True, text=True)
    if result.returncode != 0:
        print("Error listing containers:", result.stderr)
        return []
    return result.stdout.splitlines()


def remove_container(container):
    rm_result = subprocess.run(['docker', 'rm', '-f', container], capture_output=True, text=True)
    if rm_result.returncode == 0:
        print(f"Successfully removed container: {container}")
    else:
        print(f"Error removing container {container}: {rm_result.stderr}")


# remove `/var/lib/msmc/` directory if it exists
if os.path.exists('/var/lib/msmc/'):
    os.system('rm -rf /var/lib/msmc/')
    print('Removed /var/lib/msmc/ directory')

# remove all docker containers which name starts with `MSMC`:
containers = list_containers()
msmc_containers = [name for name in containers if name.startswith("MSMC")]
for container in msmc_containers:
    remove_container(container)
