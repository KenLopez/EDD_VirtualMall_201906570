import React from 'react'
import {BrowserRouter as Router, Route, Redirect} from 'react-router-dom'
import {Segment} from 'semantic-ui-react'
import Home from './components/Home'
import Reporte from './components/Reporte'
import Tienda from './components/Tienda'
import CarritoCompra from './components/CarritoCompra'
import './App.css'
import ArchivosContainer from './components/ArchivosContainer'
import Login from './components/Login'
import Registro from './components/Registro'
import TiendasContainer from './components/TiendasContainer'
import Inventario from './components/Inventario'

function App() {
  return (
    <>
      <Router>
        <Route exact path="/">
          <Redirect to="/Login" />
        </Route>
        <Route path="/Login" component={Login}/>
        <Route path="/Registro" component={Registro}/>
        <Route path="/Home" component={Home}/>
        <Route path="/CargarArchivo" component={ArchivosContainer}/>
        <Route path="/CarritoCompra" component={CarritoCompra}/>
        <Route path="/Reporte" component={Reporte}/>
        <Route exact path="/Tienda/:Departamento/:Nombre/:Calificacion" component={Tienda}/>
        <Route exact path="/Inventarios/:Departamento/:Nombre/:Calificacion" component={Inventario}/>
        <Route exact path="/Inventarios" component={TiendasContainer}/>
        <Segment inverted color='black' className="Footer">
          <center>
            Kenneth Haroldo López López<br/>201906570<br/>USAC 2021
          </center>
        </Segment>
      </Router>
    </>
  )
}

export default App
