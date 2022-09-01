import { Row, Layout, Image, Typography } from 'antd';
import myImage from '../../assets/my.jpg';

const { Title } = Typography;

function BannerComponent(): any
{
    return(
        <>
            <Layout className="banner-home">
                <Row>
                    <Image width={300} src={myImage} />
                </Row>
                <Row>
                    <Title level={1}>
                        INDEPENDENT FULL-STACK <br /> SOFTWARE ENGINNER
                    </Title>
                </Row>
            </Layout>
        </>
    );
}

export default BannerComponent;