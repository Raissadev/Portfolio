import { Layout } from 'antd';
import Head from "../components/Header";
import BannerComponent from '../components/Home/BannerComponent';

function Home(): any
{
    return(
        <>
            <Head />
            <Layout>
                <BannerComponent />
            </Layout>
        </>
    );
}

export default Home;