EXPECTED_HASH="76d6edc1050629c0be0f82543d7da36f7c27ef2e7978e28c20bb5b9e7677c7c8b485af3c6a34319f402706592a1b381ad8ec96458104b23814c871e2b4870eab"

ACTUAL_HASH=$(sha512sum $DOWNLOADS_ROOT$SPARK_ARCHIVE  | awk '{print $1}')

if [ "$ACTUAL_HASH" == "$EXPECTED_HASH" ]; 
then
    echo "OK: Hash of archive file matches expected hash"
else
    echo "ERROR: Hash of archive file does not match expected hash"
    exit 1
fi

tar -xvzf $DOWNLOADS_ROOT$SPARK_ARCHIVE -C $UNPACK_ROOT

