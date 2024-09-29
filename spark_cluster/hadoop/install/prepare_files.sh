EXPECTED_HASH="6f653c0109f97430047bd3677c50da7c8a2809d153b231794cf980b3208a6b4beff8ff1a03a01094299d459a3a37a3fe16731629987165d71f328657dbf2f24c"

ACTUAL_HASH=$(sha512sum $DOWNLOADS_ROOT$HADOOP_ARCHIVE | awk '{print $1}')

if [ "$ACTUAL_HASH" == "$EXPECTED_HASH" ]; 
then
    echo "OK: Hash of archive file matches expected hash"
else
    echo "ERROR: Hash of archive file does not match expected hash"
    exit 1
fi

tar -xvzf $DOWNLOADS_ROOT$HADOOP_ARCHIVE -C $UNPACK_ROOT