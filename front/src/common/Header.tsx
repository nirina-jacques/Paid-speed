import { Link } from 'react-router-dom'

export default function Header() {
    return (
        <div className="wrapper_border">
                <div className="header row">
                    <div className="col-sm-2 logo">
                        <h2>
                            <Link to="/">Oasis Food</Link>
                        </h2> 
                    </div>
                    <div className="col-sm-8">
                        <div className="row wrapper_cnx_reg">
                            <div className="col-sm-3 location">
                                <span>
                                    <i className="fas fa-map-marker-alt"></i>
                                </span>
                                <span>
                                    13/80 Face BFV SG Tanambao5
                                </span>
                            </div>
                            <div className="col-sm-6 search-content text-center">
                                    <div className="info_bull">
                                        <div>Livraison gratuit</div>
                                        <div>Tel: 038 73 217 67</div>
                                    </div>
                            </div>
                            <div className="col-sm-2 identifier_register d-flex">
                                <div className="cnx link">
                                    <Link to="#">Connexion</Link>
                                </div>
                                <div className="register link">
                                    <Link to="#">Inscription</Link>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="col-sm-2">&nbsp;</div>
                </div>
            </div>
    )
}