<script>
    import Top from "./Top.svelte";
    import Footer from "./Footer.svelte";
    import Login from "./Login.svelte";
    import ChooseServer from "./ChooseServer.svelte";
</script>
<style>
    #container{
        display: flex;
        flex-direction: column;
        height: 100vh;
    }
</style>

<div id="container">
    <Top/>
    <ChooseServer/>
    <Footer/>
</div>