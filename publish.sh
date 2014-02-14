#!/bin/bash
for filename in ./build/*; do
    bin="${filename##*/}"
    echo "uploading $bin"
    curl -T $filename -uhalleck45:$BINTRAY_LICENSE https://api.bintray.com/content/halleck45/OSS/bin/`semver tag`/$bin

    echo "publishing $bin"
    curl -X POST -uhalleck45:$BINTRAY_LICENSE https://api.bintray.com/content/halleck45/OSS/bin/`semver tag`/publish
done


