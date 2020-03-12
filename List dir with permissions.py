import subprocess
import schedule
import time

def getPermissions():
    subprocess.Popen(['powershell',
                      '& "C:\\Users\\WimfredPC\\Documents\\Fontys\\Infra\\Semester 2\\Programming\\Folder.ps1"'], shell=False)

schedule.every(1).minutes.do(getPermissions)

while True: 
    schedule.run_pending() 
    time.sleep(1) 

