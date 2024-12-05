param([Int32]$day) 

$day_string = 'day{0:D2}' -f $day
pushd

cd .\resources
mkdir $day_string
cd $day_string
touch "input.txt"
touch "sample.txt"
popd

pushd
cd .\days
mkdir $day_string
cd $day_string
touch "$day_string.go"
"package $day_string" | Out-File -FilePath "$day_string.go"
touch "${day_string}_test.go"
"package $day_string" | Out-File -FilePath "${day_string}_test.go"
popd
