async function getKey(profileDirectory, masterPassword) {
    const key4FilePath = path.join(profileDirectory, "key4.db");
    if (!fs.existsSync(key4FilePath)) {
        throw new Error("key4.db was not found in this profile directory.");
    }

    const masterPasswordBytes = forge.util.encodeUtf8(masterPassword || "");
    const key4File = fs.readFileSync(key4FilePath);

    const key4Db = await initSqlJs().then(function(SQL){
        return new SQL.Database(key4File);
    });

    const metaData = key4Db.exec("SELECT item1, item2 FROM metadata WHERE id = \"password\";");
    if (metaData && metaData.length && metaData[0].values && metaData[0].values.length) {
        const globalSalt = toByteString(metaData[0].values[0][0].buffer);
        const item2 = toByteString(metaData[0].values[0][1].buffer);
        const item2Asn1 = forge.asn1.fromDer(item2);
        const item2Value = pbesDecrypt(item2Asn1.value, masterPasswordBytes, globalSalt);
        if (item2Value && item2Value.data === "password-check") {
            const nssData = key4Db.exec("SELECT a11 FROM nssPrivate WHERE a11 IS NOT NULL;");
            if (nssData && nssData.length && nssData[0].values && nssData[0].values.length) {
                const a11 = toByteString(nssData[0].values[0][0].buffer);
                const a11Asn1 = forge.asn1.fromDer(a11);
                return pbesDecrypt(a11Asn1.value, masterPasswordBytes, globalSalt);
            }
        } else {
            throw new Error("Master password incorrect.");
        }
    }

    throw new Error("Not able to get key from profile directory or no passwords were found.");
}
