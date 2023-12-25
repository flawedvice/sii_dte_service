# SII DTE Service

El objetivo de este servicio es generar DTEs acorde a lo documentado por el [SII](https://www.sii.cl/factura_electronica/instructivo_emision.pdf).

## Pasos:
1. Recibir CAF y llave privada para generar timbre.
2. Asignar número de folio único a cada documento.
3. Calcular el timbre electrónico con la llave privada.
4. Generar el documento XML.
5. Firmar documento completo con un certificado digital.
6. Preparar el documento para su impresión (generar PDF).
7. Permitir la recepción y emisión de correos con documentos adjuntos.

Para obtener el CAF, uno debe dirigirse al sitio del SII según se documenta [aquí](https://www.sii.cl/servicios_online/manual_autoriacionDT.pdf).

## System design:
- DB with every uploaded CAF's information.
- Access to users' private keys.
- Module to calculate digital signature (timbre electrónico).
- Module to generate XML.
- Module to sign XML docs using digital certificate.
- Module to generate PDFs from XML data.
- 