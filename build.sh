echo "start build (linux,386)";
CGO_ENABLED=0 GOOS=linux GOARCH=386  go build -o build/fchinanet_linux_x86 fchinanet.go;
echo "complete build (linux,386)";

echo "start build (linux,amd64)";
CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o build/fchinanet_linux_x86-64 fchinanet.go;
echo "complete build (linux,amd64)";

echo "start build (linux,ARMv5)";
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -o build/fchinanet_linux_ARMv5 fchinanet.go;
echo "complete build (linux,ARMv5)";

echo "start build (linux,ARMv6)";
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o build/fchinanet_linux_ARMv6 fchinanet.go;
echo "complete build (linux,ARMv6)";

echo "start build (linux,ARMv7)";
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o build/fchinanet_linux_ARMv7 fchinanet.go;
echo "complete build (linux,ARMv7)";

echo "start build (linux,ARMv8)";
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/fchinanet_linux_ARMv8 fchinanet.go;
echo "complete build (linux,ARMv8)";

echo "start build (linux,mips32)";
CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -o build/fchinanet_linux_mips32 fchinanet.go;
echo "complete build (linux,mips32)";

echo "start build (linux,mips32le)";
CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -o build/fchinanet_linux_mips32le fchinanet.go;
echo "complete build (linux,mips32le)";

echo "start build (linux,mips64)";
CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -o build/fchinanet_linux_mips64 fchinanet.go;
echo "complete build (linux,mips64)";

echo "start build (linux,mips64le)";
CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -o build/fchinanet_linux_mips64le fchinanet.go;
echo "complete build (linux,mips64le)";

echo "start build (windows,386)";
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o build/fchinanet_win_x86.exe fchinanet.go;
echo "complete build (windows,386)";

echo "start build (windows,amd64)";
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o build/fchinanet_win_x86-64.exe fchinanet.go;
echo "complete build (windows,amd64)";

echo "start build (macos,amd64)";
CGO_ENABLED=0 GOOS=darwin go build -o build/fchinanet_mac fchinanet.go;
echo "complete build (macos,amd64)";