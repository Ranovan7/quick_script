hello = "hello world" :: [Char]

doubleIt :: Int -> Int
doubleIt x = x * 2

main :: IO()
main = do
    print(doubleIt 5)
    print(doubleIt 25)
    print(doubleIt 75)
