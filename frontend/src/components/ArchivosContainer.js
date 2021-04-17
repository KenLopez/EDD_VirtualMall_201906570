import React from 'react'
import CargaArchivo from './CargaArchivo'
import '../css/Content.css'
import NavBar from './NavBar'

function ArchivosContainer() {
    return (
        <>
        <NavBar
        activo={2}/>
        <div className="Content">
            <div className="ui segment mosaico container">
                <CargaArchivo
                        Ruta='http://localhost:3000/cargarGrafo'
                        Title='Cargar Grafo'
                />
                <br/>
                <CargaArchivo
                        Ruta='http://localhost:3000/cargartienda'
                        Title='Cargar Tiendas'
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
                    Ruta='http://localhost:3000/CargarUsuarios'
                    Title='Cargar Usuarios'
                />
            </div>
        </div>
        </>
    )
}

export default ArchivosContainer
