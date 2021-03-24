import React from 'react'
import CargaArchivo from './CargaArchivo'
import '../css/Content.css'

function ArchivosContainer() {
    return (
        <div className="Content">
            <div className="ui segment mosaico container">
                <CargaArchivo
                        Ruta='http://localhost:3000/cargartienda'
                        Title='Cargar Tienda'
                />
                <br/>
                <CargaArchivo
                    Ruta='http://localhost:3000/CargarInventarios'
                    Title='Cargar Inventarios'
                />
                <br/>
                <CargaArchivo
                    Ruta='http://localhost:3000/CargarPedidos'
                    Title='Cargar Pedidos'
                />
            </div>
        </div>
    )
}

export default ArchivosContainer
