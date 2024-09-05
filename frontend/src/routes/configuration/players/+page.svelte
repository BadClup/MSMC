<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import OptionsList from "./../optionsList.svelte";

    interface Player{
        name: string;
        isOnline: boolean;
    }

    const players: Player[] = [{name: "player1", isOnline:false}, 
    {name: "player2", isOnline:true}, 
    {name: "player3", isOnline:false},
    {name: "player4", isOnline:true}, 
    {name: "player5", isOnline:false}]
</script>
<style>
    section{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
        padding: 0;
        background-color: rgb(88, 88, 88);
        color: white;
    }

    section > div{
        width: 100%;
        flex: 1;
    }

    nav{
        display: inline-block;
        float: left;
        height: 100%;
        width: 240px;
    }

    div > section{
        width: 320px;
        height: 100%;
        float: right;
        box-sizing: border-box;
        padding-top: 35px;
    }

    section > a{
        box-sizing: border-box;
        height: 100px;
        margin-top: 60px;
        margin-right: 40px;
        padding-top: 30px;
        background-color: red;
        text-decoration: none;
        font-weight: bold;
        font-size: 30px;
        color: white;
        text-align: center;
    }

    .playersContainer{
        display: inline-block;
        padding: 3%;
        width: calc(100% - 690px)
    }

    .playerWrapper{
        box-sizing: border-box;
        display: inline-block;
        width: 100%;
        height: 50px;
        margin-top: 22px;
        background-color: rgb(48, 48, 48);
    }

    .playerInfo{
        display: inline-block;
        background-color: rgb(48, 48, 48);
        color: white;
        font-weight: 600;
        font-size: 20px;
        padding: 13px;
    }

    .playerInfo::before{
        content: url($lib/images/steve.png);
        margin-right: 10px;
    }

    .buttonsContainer{
        display: inline-block;
        float: right;
        background-color: rgb(48, 48, 48);
        color: white;
        padding: 10px;
    }

    button{
        width: 30px;
        height: 30px;
    }
</style>

<section>
    <Top showServers={true} isLoggedIn={true}/>
    <div>
        <nav><OptionsList name={"name"} version={"version"} author={"author"}/></nav>
        <div class="playersContainer">
            {#each players as player}
                <div class="playerWrapper">
                    <div class="playerInfo">{player.name}</div>
                    <div class="buttonsContainer">
                        <button>O1</button>
                        <button>O2</button>
                        <button>O3</button>
                        <button>O4</button>
                    </div>
                </div>
            {/each}
        </div>
        <section>
            <a href="#" style="background-color: green">Whitelist</a>
            <a href="#" style="background-color: green">Operators</a>
            <a href="#" style="background-color: red">Banned players</a>
            <a href="#" style="background-color: red">Banned IP's</a>
        </section>
    </div>
    <Footer/>
</section>