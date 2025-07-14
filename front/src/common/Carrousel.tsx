import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';

export default function Carrousel() {
    return (
        <div className="container">
            <div className="row">
                <div className="col-12">
                <Swiper
                    spaceBetween={50}
                    slidesPerView={3}
                    onSlideChange={() => console.log('slide change')}
                    onSwiper={(swiper) => console.log(swiper)}
                    >
                    <SwiperSlide>Slide 1</SwiperSlide>
                    <SwiperSlide>Slide 2</SwiperSlide>
                    <SwiperSlide>Slide 3</SwiperSlide>
                    <SwiperSlide>Slide 4</SwiperSlide>
                    ...
                    </Swiper>
                </div>
            </div>
        </div>    
    )
}