import React from 'react'
import { Redirect, Route, Router } from 'react-router'
import NavBar from './NavBar'
import Home from './Home'
import ArchivosContainer from './ArchivosContainer'
import CarritoCompra from './CarritoCompra'
import Reporte from './Reporte'
import Tienda from './Tienda'
import Footer from './Footer'

function Loged() {
    return (
    <>
      <Router>
        <NavBar/>
        <Route exact path="/Loged">
          <Redirect to="/Home" />
        </Route>
        <Route path="/Home" component={Home}/>
        <Route path="/CargarArchivo" component={ArchivosContainer}/>
        <Route path="/CarritoCompra" component={CarritoCompra}/>
        <Route path="/Reporte" component={Reporte}/>
        <Route exact path="/Tienda/:Departamento/:Nombre/:Calificacion" component={Tienda}/>
      </Router>
      <Footer/>
    </>
    )
}

export default Loged
