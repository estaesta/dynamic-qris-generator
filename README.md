# Project dynamic-qris-generator

An API to generate dynamic QRIS from a static QRIS

## Endpoints

### 1. **POST /api/qris/extract**

This endpoint accepts a QRIS image file and extracts the QRIS data from the image.

#### Request
- **Method**: `POST`
- **URL**: `/api/qris/extract`
- **Content-Type**: `multipart/form-data`
- **Form Data**: 
  - `qris` (required): The QRIS image file to be extracted. (Max file size: 5MB)

#### Response
- **Content-Type**: `application/json`
- **Status**: `200 OK`
- **Body**: A JSON object with the extracted QRIS data.

```json
{
  "data": "00020101021126640017ID.CO.BANKBSI.WWW0118936004510000305481021000002039790303URE51440014ID.CO.QRIS.WWW0215ID10243141810570303URE5204866153033605802ID5918SABILILLAH YAYASAN6006MALANG6105651416304522C"
}
```

---

### 2. **POST /api/qris/convert**

This endpoint accepts QRIS data as a string and generates a dynamic QRIS image based on that data.

#### Request
- **Method**: `POST`
- **URL**: `/api/qris/convert`
- **Content-Type**: `application/json`
- **Body**:
  - `data` : The QRIS string data that needs to be encoded into a QRIS image.
  - `amount` : The transaction amount associated with the QRIS data.

```json
{
  "data": "00020101021126640017ID.CO.BANKBSI.WWW0118936004510000305481021000002039790303URE51440014ID.CO.QRIS.WWW0215ID10243141810570303URE5204866153033605802ID5918SABILILLAH YAYASAN6006MALANG6105651416304522C",
  "amount": 100000
}
```

#### Response
- **Content-Type**: `image/jpeg`
- **Status**: `200 OK`
- **Body**: A raw JPEG image representing the generated QRIS.

---


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

