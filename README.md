## Lista LCO - SAT

Programa en Go (Golang) para leer, descargar y procesar la lista LCO del SAT mediante [esta URL](https://cfdisat.blob.core.windows.net/lco?restype=container&comp=list&prefix=LCO_2019-10-18).

### URL

```bash
https://cfdisat.blob.core.windows.net/lco?restype=container&comp=list&prefix=LCO_2019-10-18
```
La fecha está en formato **YYYY-MM-DD**.

## PROCESO
1. **Descarga** - Descarga el archivo con la lista de BLOBs
2. **XML** - Lee la lista de BLOBs
3. **Descarga** - Se descargan la lista de BLOBs
4. **Proceso** - Se procesa cada archivo
5. **CSV** - Crea un archivo CSV
6. **Database** - Guarda la informacion en Base de Datos

### Saludos...