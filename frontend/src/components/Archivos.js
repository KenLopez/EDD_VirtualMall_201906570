import React from 'react'
import CargarArchivo from './CargarArchivo'
import '../css/Content.css'

function Archivos() {
    return (
        <div className="Content">
            <div className="ui segment mosaico container">
                <CargarArchivo
                        Ruta='http://localhost:3000/cargartienda'
                        Title='Cargar Tienda'
                />
                <br/>
                <CargarArchivo
                    Ruta='http://localhost:3000/CargarInventarios'
                    Title='Cargar Inventarios'
                />
                <br/>
                <CargarArchivo
                    Ruta='http://localhost:3000/CargarPedidos'
                    Title='Cargar Pedidos'
                />
            </div>
        </div>
    )
}

export default Archivos
