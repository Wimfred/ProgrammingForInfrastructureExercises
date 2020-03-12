$path = 'C:\Users\WimfredPC\Documents\Fontys\Infra\'
$a = Get-ChildItem -Directory -Recurse -Force -Path $path
$i = Foreach ($b in $a){
Get-Acl  $b.FullName
} 
$date = Get-Date -UFormat "%d-%m-%Y"
$time = Get-Date -Format "HH mm ss"
$fulldate = "$date($time)"
$filename = "list_$fulldate.txt"
$i > "C:\Users\WimfredPC\Documents\Fontys\Infra\Semester 2\Programming\Output\$filename"


