<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: Tree.Service.proto

namespace Tree;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 *搜索响应体
 *
 * Generated from protobuf message <code>tree.SearchResponse</code>
 */
class SearchResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .tree.SearchResult ret = 1;</code>
     */
    private $ret;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Tree\SearchResult[]|\Google\Protobuf\Internal\RepeatedField $ret
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\TreeService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .tree.SearchResult ret = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getRet()
    {
        return $this->ret;
    }

    /**
     * Generated from protobuf field <code>repeated .tree.SearchResult ret = 1;</code>
     * @param \Tree\SearchResult[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setRet($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Tree\SearchResult::class);
        $this->ret = $arr;

        return $this;
    }

}

