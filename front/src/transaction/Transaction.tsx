import { useState, useEffect } from "react"

import Select, {components} from 'react-select'
import Header from "../common/Header";
import Footer from "../common/Footer";

import ImgOrange from "../../public/image/orange2.png"
import ImgMvola from "../../public/image/mvola.png"
import ImgAirtel from "../../public/image/airtel.png"

type Operator = {
    value: number;
    label: string;
    icon: string;
}

export default function Transaction() {
    const [operatorMobile, setoperatorMobile] = useState<Operator[]>([])
    useEffect(() => {
        setoperatorMobile([
            {
                value:1,
                label: "Mvola",
                icon: ""
            },
            {
                value:2,
                label: "Orange Money",
                icon: ImgOrange
            },
            {
                value: 3,
                label: "Airtel Money",
                icon: ""
            }
        ])
    }, [])
    const { Option } = components;
    const IconOption = (props: any) => (
        <Option {...props}>
          <div className="d-flex">
            <div className="col-sm-8" style={{textAlign:"left"}}>{props.data.label}</div>
            <div className="col-sm-4">
                {
                    props.data.value == 1 ? 
                    <img src={ImgMvola} style={{ width: 36}} /> 
                    : props.data.value == 2 
                    ? <img src={ImgOrange} style={{ width: 36}} /> 
                    : <img src={ImgAirtel} style={{ width: 36}} />
                }
            </div>
          </div>
        </Option>
      );
    return (
        <div>
            <Header></Header>
            <div className="wrap_transfert">
                <div className="content-transfert">
                    <div>Détails transfert</div>
                    <div>
                        <div className="email-paypal form-group">
                            <input type="text" className="form-control"  id="email_paypal" placeholder="Email Paypal" />
                        </div>
                        <div className="row">
                            <div className="col-sm-3 form-group space_email">
                                <Select
                                    options={operatorMobile}
                                    components={{ Option: IconOption }}
                                />
                            </div>
                            <div className="col-sm-4 form-group space_email">
                                <input type="text" className="form-control"  id="amount" placeholder="Somme Ar" />
                            </div>
                            <div className="col-sm-4 form-group">
                                <input type="text" className="form-control"  id="mobile" placeholder="Numéro téléphone" />
                            </div>
                        </div>
                    </div>
                    <button className="btn btn-primary transaction_btn">Transfert</button>
                </div>
            </div>
            <Footer></Footer>
        </div>
    )
}