set GOARCH=386
set GOOS=windows
rm .\resources\dummyprogram\dummyprogram.exe
go build -o resources\dummyprogram\dummyprogram.exe resources\dummyprogram\main.go
rm .\resources\NaAVFakeProgramSpawner\NaAVFakeProgramSpawner.exe
go build -o resources\NaAVFakeProgramSpawner\NaAVFakeProgramSpawner.exe resources\NaAVFakeProgramSpawner\main.go