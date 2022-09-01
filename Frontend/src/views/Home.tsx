import { Layout } from 'antd';
import Head from "../components/Header";
import BannerComponent from '../components/Home/BannerComponent';
import BoxComponent from '../components/Home/BoxComponent';
import MessageComponent from '../components/Home/MessageComponent';

function Home(): any
{
    return(
        <>
            <Head />

            <Layout>
                <BannerComponent />
            </Layout>

            <Layout>
                <MessageComponent />
            </Layout>

            <Layout>
                <BoxComponent />
            </Layout>
        </>
    );
}

export default Home;