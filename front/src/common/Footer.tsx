
export default function Footer() {
    return (
        <footer className="text-center footer">
            <div className="container p-4">
                <section className="mb-4">
                    <a className="btn btn-outline-light btn-floating m-1" href="https://www.facebook.com/oasis.food.tanambao5" role="button">
                        <i className="fab fa-facebook"></i>
                    </a>
                    <a className="btn btn-outline-light btn-floating m-1" href="#!" role="button">
                        <i className="fab fa-instagram"></i>
                    </a>
                    <a className="btn btn-outline-light btn-floating m-1" href="#!" role="button">
                        <i className="fab fa-linkedin-in"></i>
                    </a>
                </section>
            </div>
            <div className="oasis-color">
                <span className="phone_txt">
                    Tel: +261 38 73 217 67
                </span>
            </div>
            <div className="text-center p-3 copy_rigth">
                <a className="oasis-color" href="https://oasisfood.vercel.app/">Oasis Food</a>
            </div>
        </footer>
    )
}