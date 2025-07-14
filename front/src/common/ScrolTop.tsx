import { useState, useEffect } from 'react';

export default function ScrolTop() {
    const [scrollY, setScrollY] = useState(0)
    useEffect(() => {
        const onScroll = () => setScrollY(window.scrollY);
        window.addEventListener('scroll', onScroll);
        return () => window.removeEventListener('scroll', onScroll);
      }, []);

    const ScrolTo = function() {
        window.scrollTo(0, 0)
    }
    return (
        <div>
            <button onClick={ScrolTo} className={scrollY > 588 ? "scroll": "d-none"}>
                <i className="fas fa-chevron-up"></i>
            </button>
        </div>
    )
}