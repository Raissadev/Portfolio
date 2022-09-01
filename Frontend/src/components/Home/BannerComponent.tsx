import { Row, Col, Layout, Image, Typography, Button } from 'antd';
import { RightOutlined, CaretLeftFilled } from '@ant-design/icons';
import myImage from '../../assets/my.jpg';

const { Title } = Typography;

function BannerComponent(): any
{
    return(
        <>
            <Layout className="banner-home" id="home">
                <Row>
                    <Image width={300} src={myImage} />
                </Row>
                <Row>
                    <Col span={24}>
                        <Title level={1}>
                            INDEPENDENT FULL-STACK <br /> SOFTWARE ENGINNER
                        </Title>
                    </Col>
                    <Col span={24}>
                        <Button type="primary" size="large">
                            Send mail
                            <RightOutlined />
                        </Button>
                    </Col>
                </Row>
                <Row className="attach">
                    <Title level={4}>
                        <span>--------------- <CaretLeftFilled /></span>
                        SCROLL TO SEE MORE
                    </Title>
                </Row>
            </Layout>
        </>
    );
}

export default BannerComponent;