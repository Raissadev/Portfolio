import { Row, Col, Layout, Typography, Button } from 'antd';
import { RightOutlined } from '@ant-design/icons';

const { Title } = Typography;
const { Footer } = Layout;

function Foot(): any
{
    return(
        <>
            <Footer>
                <Row align="middle" justify="center">
                    <Col span={24}>
                        <Title level={2}>
                            Want to create <br/> amazing stuff
                        </Title>
                    </Col>
                    <Col span={24}>
                        <Button type="primary" size="large">
                            Send mail
                            <RightOutlined />
                        </Button>
                    </Col>
                </Row>
                <Row align="middle" justify="center" id="medias">
                    <Button href="https://github.com/Raissadev" type="link">GitHub</Button>
                    <Button href="https://www.linkedin.com/in/raissadev" type="link">Linkedin</Button>
                    <Button href="https://www.instagram.com/raissa_dev" type="link">Instagram</Button>
                    <Button href="https://twitter.com/Raissa_Dev" type="link">Twitter</Button>
                </Row>
                <Row align="middle" justify="center">
                    <Title level={4}>
                        ©2022 raissadev - All rights reserved
                    </Title>
                </Row>
            </Footer>
        </>
    );
}

export default Foot;