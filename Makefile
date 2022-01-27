build:
	go build -trimpath -buildmode=plugin -o ./bin/librcloneplugin_backend_mwe.so .
	cp -f ./bin/librcloneplugin_backend_mwe.so ${RCLONE_PLUGIN_PATH}/librcloneplugin_backend_mwe.so

clean:
	rm -rf ./bin