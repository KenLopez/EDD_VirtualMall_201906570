import React from 'react'
import {BrowserRouter as Router, Route, Redirect} from 'react-router-dom'
import {Segment} from 'semantic-ui-react'
import Home from './components/Home'
import Reporte from './components/Reporte'
import NavBar from './components/NavBar'
import Tienda from './components/Tienda'
import CarritoCompra from './components/CarritoCompra'
import './App.css'
import Archivos from './components/Archivos'

function App() {
  return (
    <>
      <Router>
        <NavBar/>
        <Route exact path="/">
          <Redirect to="/Home" />
        </Route>
        <Route path="/Home" component={Home}/>
        <Route path="/CargarArchivo" component={Archivos}/>
        <Route path="/CarritoCompra" component={CarritoCompra}/>
        <Route path="/Reporte" component={Reporte}/>
        <Route exact path="/Tienda/:Departamento/:Nombre/:Calificacion" component={Tienda}/>
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
