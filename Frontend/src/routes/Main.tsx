import React from "react";
import { Routes, Route } from "react-router-dom";

import Home from "../views/Home"

import "../styles/App.less"

export default function Main(): any
{
    return(
        <>
            <Routes>
                <Route path="/" element={ <Home /> } />
            </Routes>
        </>
    )
}