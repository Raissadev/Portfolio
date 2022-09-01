import { Layout } from 'antd';
import Head from "../components/Header";
import BannerComponent from '../components/Home/BannerComponent';
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
        </>
    );
}

export default Home;