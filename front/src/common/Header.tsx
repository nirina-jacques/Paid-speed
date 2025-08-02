import { Link } from 'react-router-dom'

export default function Header() {
    return (
        <div className="wrapper_border">
            <div className="header row">
                <div className="col-sm-2 logo">
                    <h2>
                        <Link to="/">Paypal Speed</Link>
                    </h2> 
                </div>
                <div className="col-sm-8">
                    <div className="row wrapper_cnx_reg">
                        <div className="col-sm-10 search-content d-flex">
                        <div className="col-sm-8">
                                <Link to="/history">Mes transaction</Link>
                            </div>
                            <div className="col-sm-2 link-to">
                                <Link to="/">Nous contacter</Link>
                            </div>
                            <div className="col-sm-2 link-to">
                                <Link to="/">FAQ</Link>
                            </div>
                        </div>
                        <div className="col-sm-2 identifier_register d-flex">
                            <div className="cnx link">
                                <Link to="/login">Connexion</Link>
                            </div>
                            <div className="register link">
                                <Link to="/register">Inscription</Link>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="col-sm-2">&nbsp;</div>
            </div>
        </div>
    )
}