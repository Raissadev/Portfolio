import { Row, Col, Layout, Typography, Divider, Image } from 'antd';
import networks from '../../assets/networks.jpg';
import back from '../../assets/back.jpg';
import front from '../../assets/front.jpg';

const { Title } = Typography;

function InterestsComponent(): any
{
    return(
        <>
            <Layout className="interests-component default">
                <Row>
                    <Col span={24}>
                        <Divider orientation="left" orientationMargin="0">
                            <Title level={2}>
                                DevOps
                            </Title>
                        </Divider>
                    </Col>
                    <Col>
                        <Image width={1000} src={networks}/>
                    </Col>
                </Row>
                <Row>
                    <Col span={24}>
                        <Divider orientation="right" orientationMargin="0">
                            <Title level={2}>
                                Backend
                            </Title>
                        </Divider>
                    </Col>
                    <Col>
                        <Image width={1000} src={back}/>
                    </Col>
                </Row>
                <Row>
                    <Col span={24}>
                        <Divider orientation="left" orientationMargin="0">
                            <Title level={2}>
                                Frontend
                            </Title>
                        </Divider>
                    </Col>
                    <Col>
                        <Image width={1000} src={front}/>
                    </Col>
                </Row>
            </Layout>
        </>
    );
}

export default InterestsComponent;