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
                <br/>
                <CargaArchivo
                    Ruta='http://localhost:3000/CargarPedidos'
                    Title='Cargar Usuarios'
                />
            </div>
        </div>
    )
}

export default ArchivosContainer
