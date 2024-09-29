EXPECTED_HASH="3b041e538282c15ad3b0ab2f05d558e1be2d16e047c2a98ca9c08176c4a50e7101bd2199aec97a66fdabf4b3f490c2c1cefd547bbe35a1b3e3143e066948ee6d"

ACTUAL_HASH=$(sha512sum $DOWNLOADS_ROOT$SPARK_ARCHIVE  | awk '{print $1}')

if [ "$ACTUAL_HASH" == "$EXPECTED_HASH" ]; 
then
    echo "OK: Hash of archive file matches expected hash"
else
    echo "ERROR: Hash of archive file does not match expected hash"
    exit 1
fi

tar -xvzf $DOWNLOADS_ROOT$SPARK_ARCHIVE -C $UNPACK_ROOT

